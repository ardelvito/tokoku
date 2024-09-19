import { Box, Flex, Text } from "@chakra-ui/react";
import BottomBar from "../Components/BottomBar";

const Transactions = () => {
  return (
    <Flex direction="column" minH="100vh" justifyContent="space-between">
      {/* Konten halaman */}
      <Box flex="1" p={4}>
        <Text fontSize="2xl" fontWeight="bold">
          Transactions
        </Text>
      </Box>
      
      <BottomBar />
    </Flex>
  );
};

export default Transactions;
