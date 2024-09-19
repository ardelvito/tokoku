import { Box, Flex, Text } from "@chakra-ui/react";
import BottomBar from "../Components/BottomBar";

const Cart = () => {
  return (
    <Flex direction="column" minH="100vh" justifyContent="space-between">
      {/* Konten halaman */}
      <Box flex="1" p={4}>
        <Text fontSize="2xl" fontWeight="bold">
          Cart
        </Text>
      </Box>

      <BottomBar />
    </Flex>
  );
};

export default Cart;
