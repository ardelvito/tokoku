import { Flex, IconButton } from "@chakra-ui/react";
import { useNavigate, useLocation } from "react-router-dom";
import { AiOutlineHome, AiOutlineUser , AiOutlineShoppingCart, AiOutlineTransaction } from "react-icons/ai";
import {jwtDecode} from "jwt-decode";

const BottomBar = () => {
  const navigate = useNavigate();
  const location = useLocation();

  // Check the current path for highlighting the active route
  const isActive = (path) => location.pathname === path;

  const token = localStorage.getItem("token");
  let userId;

    if (token) {
    try {
      const decoded = jwtDecode(token);
      userId = decoded.user_id; // Mendapatkan userId dari token
      console.log(userId)
    } catch (err) {
      console.error("Failed to decode token:", err);
    }
  }

  return (
    <Flex
      as="nav"
      justifyContent="space-around"
      p={4}
      bg="green.500"  // Green background color for the bar
      boxShadow="md"
    >
      <IconButton
        aria-label="Home"
        icon={<AiOutlineHome />}
        onClick={() => navigate("/home")}
        color={isActive("/home") ? "yellow.300" : "white"} // Highlight active icon
        bg={isActive("/home") ? "green.700" : "green.500"} // Highlight background for active
        _hover={{ bg: "green.600" }} // Hover effect
      />
      <IconButton
        aria-label="Transactions"
        icon={<AiOutlineTransaction />}
        onClick={() => navigate("/transactions")}
        color={isActive("/transactions") ? "yellow.300" : "white"} // Highlight active icon
        bg={isActive("/transactions") ? "green.700" : "green.500"} // Highlight background for active
        _hover={{ bg: "green.600" }} // Hover effect
      />
      <IconButton
        aria-label="Cart"
        icon={<AiOutlineShoppingCart />}
        onClick={() => navigate("/cart")}
        color={isActive("/cart") ? "yellow.300" : "white"} // Highlight active icon
        bg={isActive("/cart") ? "green.700" : "green.500"} // Highlight background for active
        _hover={{ bg: "green.600" }} // Hover effect
      />
      <IconButton
        aria-label="Profile"
        icon={<AiOutlineUser  />}
        onClick={() =>  navigate(`/profile/${userId}`)}
        color={isActive(`/profile/${userId}`) || isActive(`/editprofile`) ? "yellow.300" : "white"} // Highlight active icon
        bg={isActive(`/profile/${userId}`) || isActive(`/editprofile`) ? "green.700" : "green.500"} // Highlight background for active
        _hover={{ bg: "green.600" }} // Hover effect
      />
    </Flex>
  );
};

export default BottomBar;
