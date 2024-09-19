import { Box, Button, FormControl, FormLabel, Input, VStack, Heading, Text, Link, Center, Stack, Alert, AlertIcon } from "@chakra-ui/react";
import { useState } from "react";
import axios from "axios";
import { Link as RouterLink, useNavigate } from "react-router-dom";

const Register = () => {
  const [formData, setFormData] = useState({
    email: "",
    password: "",
    name: "",
    phone: "",
    address: "",
  });
  const [error, setError] = useState("");
  const [successMessage, setSuccessMessage] = useState("");
  const navigate = useNavigate(); // Hook untuk memindahkan route

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError(""); // Reset error sebelum submit
    try {
      const response = await axios.post("http://localhost:8080/register", formData);

      // Jika pendaftaran berhasil, tampilkan pesan sukses dan pindahkan ke halaman login setelah beberapa detik
      if (response.data.message === "Registration successful") {
        setSuccessMessage(response.data.message);

        // Tunggu beberapa detik sebelum memindahkan ke halaman login
        setTimeout(() => {
          navigate("/login");
        }, 2000); // 2 detik
      }
    } catch (err) {
      setError("Failed to register. Please try again.");
    }
  };

  return (
    <Center minH="100vh" bg="gray.50">
      <VStack spacing={8} w="full" maxW="md" p={6} alignItems="center">
        {/* Judul TOKOKU */}
        <Heading as="h1" size="2xl" color="green.500">
          TOKOKU
        </Heading>

        {/* Card form register */}
        <Box
          w="full"
          p={8}
          borderWidth={1}
          borderRadius="lg"
          boxShadow="lg"
          bg="white"
        >
          <form onSubmit={handleSubmit}>
            <VStack spacing={4} alignItems="flex-start">
              <FormControl id="email" isRequired>
                <FormLabel>Email</FormLabel>
                <Input name="email" type="email" onChange={handleChange} />
              </FormControl>

              <FormControl id="password" isRequired>
                <FormLabel>Password</FormLabel>
                <Input name="password" type="password" onChange={handleChange} />
              </FormControl>

              <FormControl id="name" isRequired>
                <FormLabel>Name</FormLabel>
                <Input name="name" onChange={handleChange} />
              </FormControl>

              <FormControl id="phone" isRequired>
                <FormLabel>Phone</FormLabel>
                <Input name="phone" onChange={handleChange} />
              </FormControl>

              <FormControl id="address" isRequired>
                <FormLabel>Address</FormLabel>
                <Input name="address" onChange={handleChange} />
              </FormControl>

              {error && <Box color="red.500">{error}</Box>}

              {/* Tampilkan pesan sukses jika registrasi berhasil */}
              {successMessage && (
                <Alert status="success">
                  <AlertIcon />
                  {successMessage}
                </Alert>
              )}

              <Button type="submit" colorScheme="green" w="full" mt={4}>
                Register
              </Button>
            </VStack>
          </form>
          {/* Teks login di bawah form */}
          <Stack direction="row" justifyContent="center" mt={4}>
            <Text fontSize="sm">
              Already have an account?{" "}
              <Link as={RouterLink} to="/login" color="green.500">
                Login
              </Link>
            </Text>
          </Stack>
        </Box>
      </VStack>
    </Center>
  );
};

export default Register;
