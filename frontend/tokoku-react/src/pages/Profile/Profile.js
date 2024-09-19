import { Box, Flex, Text, Spinner, Button } from "@chakra-ui/react";
import { AiOutlineEdit, AiOutlineUser, AiOutlineMail, AiOutlinePhone, AiOutlineHome } from "react-icons/ai";
import BottomBar from "../../Components/BottomBar";
import { useEffect, useState } from "react";
import { jwtDecode } from "jwt-decode";
import { useParams, useNavigate } from "react-router-dom";
import axios from "axios";

const Profile = () => {
  const { id } = useParams();
  const [userInfo, setUserInfo] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      try {
        const decoded = jwtDecode(token);
        const userId = decoded.user_id;

        if (userId !== Number(id)) {
          setError("Unauthorized access");
          setLoading(false);
          return;
        }

        const fetchUserData = async () => {
          try {
            const response = await axios.get(`http://localhost:8080/profile/${userId}`, {
              headers: {
                Authorization: `Bearer ${token}`,
              },
            });
            setUserInfo(response.data);
          } catch (err) {
            setError("Error fetching user data");
          } finally {
            setLoading(false);
          }
        };

        fetchUserData();
      } catch (err) {
        setError("Token is invalid");
        setLoading(false);
      }
    } else {
      setError("No token found");
      setLoading(false);
    }
  }, [id]);

const handleStoreButtonClick = () => {
    if (userInfo.StoreStatus === true) {
        // Pastikan Store tidak undefined
        if (userInfo.store) { // Perhatikan bahwa 'store' dengan huruf kecil
            navigate(`/myproducts/${userInfo.store.id}`); // Menggunakan id toko
        } else {
            console.error("Store information is not available");
            setError("Store information is not available.");
        }
    } else {
        navigate('/createstore');
    }
};


  const handleEditProfileClick = () => {
    navigate('/editprofile');
  };

  const handleLogout = () => {
    localStorage.removeItem("token");
    navigate('/login');
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
    <Flex direction="column" minH="100vh" justifyContent="space-between">
      <Box flex="1" p={6} boxShadow="md" borderRadius="lg" bg="white" maxW="600px" mx="auto" mt={6}>
        <Text fontSize="3xl" fontWeight="bold" mb={6} textAlign="center">
          Profile
        </Text>
        <Flex direction="column" gap={4}>
          <Flex align="center">
            <AiOutlineUser mr={2} />
            <Text fontSize="lg">Name: {userInfo.Name}</Text>
          </Flex>
          <Flex align="center">
            <AiOutlineMail mr={2} />
            <Text fontSize="lg">Email: {userInfo.Email}</Text>
          </Flex>
          <Flex align="center">
            <AiOutlinePhone mr={2} />
            <Text fontSize="lg">Phone: {userInfo.Phone}</Text>
          </Flex>
          <Flex align="center">
            <AiOutlineHome mr={2} />
            <Text fontSize="lg">Address: {userInfo.Address}</Text>
          </Flex>
        </Flex>
        <Button
          mt={8}
          colorScheme={userInfo.StoreStatus ? "teal" : "blue"}
          onClick={handleStoreButtonClick}
          width="full"
        >
          {userInfo.StoreStatus ? "Dashboard Toko" : "Buat Toko"}
        </Button>
        <Button
          mt={4}
          colorScheme="gray"
          leftIcon={<AiOutlineEdit />}
          onClick={handleEditProfileClick}
          width="full"
        >
          Edit Profile
        </Button>
        <Button
          mt={4}
          colorScheme="red"
          onClick={handleLogout}
          width="full"
        >
          Log Out
        </Button>
      </Box>
      <BottomBar />
    </Flex>
  );
};

export default Profile;
