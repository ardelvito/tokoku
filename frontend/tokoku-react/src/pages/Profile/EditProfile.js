import { Box, Button, Flex, FormControl, FormLabel, Input, useToast } from "@chakra-ui/react";
import { useState, useEffect } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import BottomBar from "../../Components/BottomBar";
import { jwtDecode } from "jwt-decode";

const EditProfile = () => {
  const [formData, setFormData] = useState({
    Name: "",
    Phone: "",
    Address: "",
    Email: "", // Tambahkan email ke dalam state
  });
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [userId, setUserId] = useState(null);
  const toast = useToast();
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");
    const email = localStorage.getItem("email"); // Ambil email dari local storage

    if (token) {
      try {
        const decoded = jwtDecode(token);
        const fetchedUserId = decoded.user_id;
        setUserId(fetchedUserId);

        const fetchUserData = async () => {
          try {
            const response = await axios.get(`http://localhost:8080/profile/${fetchedUserId}`, {
              headers: {
                Authorization: `Bearer ${token}`,
              },
            });
            setFormData({
              Name: response.data.Name,
              Phone: response.data.Phone,
              Address: response.data.Address,
              Email: email || "", // Set email ke formData
            });
          } catch (err) {
            setError("Error fetching user data");
          }
        };

        fetchUserData();
      } catch (err) {
        setError("Token is invalid");
      }
    } else {
      setError("No token found");
    }
  }, []);

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

    const token = localStorage.getItem("token");
    
    // Log data yang akan dikirim
    console.log("Submitting data:", { ...formData });

    try {
      await axios.post(
        "http://localhost:8080/editprofile",
        { ...formData }, // Kirim data form termasuk email
        {
          headers: {
            Authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
          },
        }
      );
      toast({
        title: "Profile updated.",
        status: "success",
        duration: 3000,
        isClosable: true,
      });
      navigate(`/profile/${userId}`);
    } catch (err) {
      toast({
        title: "Failed to update profile.",
        description: "Something went wrong.",
        status: "error",
        duration: 3000,
        isClosable: true,
      });
    } finally {
      setLoading(false);
    }
  };

  if (error) {
    return <p>{error}</p>;
  }

  return (
    <Flex direction="column" minH="100vh" justifyContent="space-between">
      <Box flex="1" p={6} boxShadow="md" borderRadius="lg" bg="white" maxW="400px" mx="auto" mt={6}>
        <form onSubmit={handleSubmit}>
          <FormControl id="name" isRequired>
            <FormLabel>Name</FormLabel>
            <Input
              type="text"
              name="Name"
              value={formData.Name}
              onChange={handleInputChange}
            />
          </FormControl>
          <FormControl id="phone" isRequired mt={4}>
            <FormLabel>Phone</FormLabel>
            <Input
              type="text"
              name="Phone"
              value={formData.Phone}
              onChange={handleInputChange}
            />
          </FormControl>
          <FormControl id="address" isRequired mt={4}>
            <FormLabel>Address</FormLabel>
            <Input
              type="text"
              name="Address"
              value={formData.Address}
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
            Update Profile
          </Button>
        </form>
      </Box>

      <BottomBar />
    </Flex>
  );
};

export default EditProfile;
