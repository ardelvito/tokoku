import { Box, Button, FormControl, FormLabel, Input, VStack, Heading, Text, Link, Center, Stack, Alert, AlertIcon } from "@chakra-ui/react";
import { useState } from "react";
import axios from "axios";
import { Link as RouterLink, useNavigate } from "react-router-dom";

const Login = () => {
  const [formData, setFormData] = useState({
    email: "",
    password: "",
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
    setSuccessMessage(""); // Reset pesan sukses sebelum submit
    try {
      const response = await axios.post("http://localhost:8080/login", formData);

      // Jika login berhasil, simpan token di local storage dan pindahkan ke halaman home
      if (response.data.token) {
        setSuccessMessage("Login successful");
        localStorage.setItem("token", response.data.token); // Simpan token ke local storage

         // Simpan email ke local storage
        localStorage.setItem("email", formData.email); // Menyimpan email

        // Tunggu beberapa detik sebelum memindahkan ke halaman home
        setTimeout(() => {
          navigate("/home");
        }, 2000); // 2 detik
      }
    } catch (err) {
      // Jika ada pesan error dari server, tampilkan pesan gagal login
      if (err.response && err.response.data.error === "Invalid credentials") {
        setError("Invalid credentials. Please try again.");
      } else {
        setError("Failed to login. Please try again.");
      }
    }
  };

  return (
    <Center minH="100vh" bg="gray.50">
      <VStack spacing={8} w="full" maxW="md" p={6} alignItems="center">
        {/* Judul TOKOKU */}
        <Heading as="h1" size="2xl" color="green.500">
          TOKOKU
        </Heading>

        {/* Card form login */}
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

              {error && (
                <Alert status="error">
                  <AlertIcon />
                  {error}
                </Alert>
              )}

              {/* Tampilkan pesan sukses jika login berhasil */}
              {successMessage && (
                <Alert status="success">
                  <AlertIcon />
                  {successMessage}
                </Alert>
              )}

              <Button type="submit" colorScheme="green" w="full" mt={4}>
                Login
              </Button>
            </VStack>
          </form>

          {/* Teks register di bawah form */}
          <Stack direction="row" justifyContent="center" mt={4}>
            <Text fontSize="sm">
              Don't have an account?{" "}
              <Link as={RouterLink} to="/register" color="green.500">
                Register
              </Link>
            </Text>
          </Stack>
        </Box>
      </VStack>
    </Center>
  );
};

export default Login;
