import { Box, Button, Flex, Text, Spinner, SimpleGrid, useToast } from "@chakra-ui/react";
import BottomBar from "../Components/BottomBar";
import { useEffect, useState } from "react";
import axios from "axios";

const Home = () => {
  const [products, setProducts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const toast = useToast();

//   useEffect(() => {
//     const fetchProducts = async () => {
//       try {
//         const response = await axios.get("http://localhost:8082/allproducts/");
//         setProducts(response.data);
//       } catch (err) {
//         setError("Error fetching products");
//         toast({
//           title: "Failed to fetch products.",
//           description: "Please try again later.",
//           status: "error",
//           duration: 3000,
//           isClosable: true,
//         });
//       } finally {
//         setLoading(false);
//       }
//     };

//     fetchProducts();
//   }, []);

//   if (loading) {
//     return (
//       <Flex justifyContent="center" alignItems="center" minH="100vh">
//         <Spinner size="xl" />
//       </Flex>
//     );
//   }

//   if (error) {
//     return (
//       <Flex justifyContent="center" alignItems="center" minH="100vh">
//         <Text color="red.500">{error}</Text>
//       </Flex>
//     );
//   }

  return (
    <Flex direction="column" p={6}>
      <Text fontSize="2xl" fontWeight="bold" mb={4}>
        All Products
      </Text>
      <SimpleGrid columns={[1, 2, 3]} spacing={4}>
        {products.map((product) => (
          <Box key={product.id} p={4} borderWidth={1} borderRadius="lg" boxShadow="md">
            <Text fontSize="xl">{product.name}</Text>
            <Text mt={2}>Price: ${product.price}</Text>
            <Flex mt={4} justifyContent="space-between">
              <Button onClick={() => console.log(`View ${product.name}`)}>View</Button>
              <Button colorScheme="blue" onClick={() => console.log(`Edit ${product.name}`)}>Edit</Button>
            </Flex>
          </Box>
        ))}
      </SimpleGrid>

      {/* Bottom navigation bar */}
      <BottomBar />
    </Flex>
  );
};

export default Home;
