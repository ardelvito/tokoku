import { Box, Button, FormControl, FormLabel, Input, useToast, Flex } from "@chakra-ui/react";
import { useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import BottomBar from "../../Components/BottomBar";

const CreateProduct = () => {
  const [formData, setFormData] = useState({
    product_name: "",
    price: "",
    description: "",
    url_image: "",
  });
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const toast = useToast();
  const navigate = useNavigate();

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError(null); // Reset error sebelum submit

    const token = localStorage.getItem("token");
    const storeId = 5; // Ganti ini dengan cara yang tepat untuk mendapatkan ID toko

    // Debugging data yang akan dikirim
    console.log("Submitting data:", { ...formData });

    try {
      await axios.post(
        `http://localhost:8080/stores/${storeId}/products`,
        {
          product_name: formData.product_name,
          price: parseFloat(formData.price),
          description: formData.description,
          url_image: formData.url_image,
        },
        {
          headers: {
            Authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
          },
        }
      );
      toast({
        title: "Product created.",
        status: "success",
        duration: 3000,
        isClosable: true,
      });
      navigate(`/dashboardstore`); // Navigasi kembali ke dashboard toko
    } catch (err) {
      setError("Failed to create product. Please try again.");
      toast({
        title: "Error.",
        description: err.response?.data?.error || "Something went wrong.",
        status: "error",
        duration: 3000,
        isClosable: true,
      });
    } finally {
      setLoading(false);
    }
  };

  return (
    <Flex direction="column" minH="100vh" justifyContent="space-between">
      <Box flex="1" p={6} boxShadow="md" borderRadius="lg" bg="white" maxW="400px" mx="auto" mt={6}>
        <form onSubmit={handleSubmit}>
          <FormControl id="product_name" isRequired>
            <FormLabel>Product Name</FormLabel>
            <Input
              type="text"
              name="product_name"
              value={formData.product_name}
              onChange={handleInputChange}
            />
          </FormControl>
          <FormControl id="price" isRequired mt={4}>
            <FormLabel>Price</FormLabel>
            <Input
              type="number"
              name="price"
              value={formData.price}
              onChange={handleInputChange}
            />
          </FormControl>
          <FormControl id="description" isRequired mt={4}>
            <FormLabel>Description</FormLabel>
            <Input
              type="text"
              name="description"
              value={formData.description}
              onChange={handleInputChange}
            />
          </FormControl>
          <FormControl id="url_image" isRequired mt={4}>
            <FormLabel>Image URL</FormLabel>
            <Input
              type="text"
              name="url_image"
              value={formData.url_image}
              onChange={handleInputChange}
            />
          </FormControl>
          <Button
            mt={6}
            colorScheme="teal"
            isLoading={loading}
            type="submit"
            width="full"
          >
            Create Product
          </Button>
        </form>
      </Box>
      <BottomBar />
    </Flex>
  );
};

export default CreateProduct;
