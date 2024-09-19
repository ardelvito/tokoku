import { Box, Button, Flex, Text, Spinner, SimpleGrid, useToast, Image } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import BottomBar from "../../Components/BottomBar";

const DashboardStore = () => {
  const [products, setProducts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const toast = useToast();
  const navigate = useNavigate();

  const storeId = 5; // Ganti ini dengan cara yang tepat untuk mendapatkan ID toko

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) {
      setError("No token found");
      setLoading(false);
      return;
    }

    const fetchProducts = async () => {
      try {
        const response = await axios.get(`http://localhost:8080/myproducts/${storeId}`, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });
        setProducts(response.data);
      } catch (err) {
        setError("Error fetching products");
      } finally {
        setLoading(false);
      }
    };

    fetchProducts();
  }, [storeId]);

  const handleDelete = async (productId) => {
    const token = localStorage.getItem("token");
    try {
      await axios.delete(`http://localhost:8080/products/${productId}`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      setProducts(products.filter(product => product.id !== productId));
      toast({
        title: "Product deleted.",
        status: "success",
        duration: 3000,
        isClosable: true,
      });
    } catch (err) {
      toast({
        title: "Failed to delete product.",
        description: "Something went wrong.",
        status: "error",
        duration: 3000,
        isClosable: true,
      });
    }
  };

  if (loading) {
    return (
      <Flex justifyContent="center" alignItems="center" minH="100vh">
        <Spinner size="xl" />
      </Flex>
    );
  }

  if (error) {
    return (
      <Flex justifyContent="center" alignItems="center" minH="100vh">
        <Text color="red.500">{error}</Text>
      </Flex>
    );
  }

  return (
    <Flex direction="column" minH="100vh">
      <Box flex="1" p={6} boxShadow="md" borderRadius="lg" bg="white" w="80%" mx="auto" mt={6}>
        <Flex justifyContent="space-between" alignItems="center" mb={4}>
          <Text fontSize="3xl" fontWeight="bold">
            Daftar Produk
          </Text>
          <Button colorScheme="teal" onClick={() => navigate('/createproduct')}>
            Buat Produk
          </Button>
        </Flex>
        <SimpleGrid columns={[1, 2, 3]} spacing={4}>
          {products.map((product) => (
            <Box key={product.id} p={4} borderWidth={1} borderRadius="lg" boxShadow="md">
              <Image src={product.url_image} alt={product.product_name} borderRadius="md" mb={2} />
              <Text fontSize="xl">{product.product_name}</Text>
              <Text mt={2}>Price: ${product.price}</Text>
              <Text mt={2}>{product.description}</Text>
              <Flex mt={4} justifyContent="space-between">
                <Button onClick={() => navigate(`/products/${product.id}`)}>View</Button>
                <Button colorScheme="blue" onClick={() => navigate(`/editproduct/${product.id}`)}>Edit</Button>
                <Button colorScheme="red" onClick={() => handleDelete(product.id)}>Delete</Button>
              </Flex>
            </Box>
          ))}
        </SimpleGrid>
      </Box>
      <BottomBar />
    </Flex>
  );
};

export default DashboardStore;
