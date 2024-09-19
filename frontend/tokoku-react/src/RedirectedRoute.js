import { Navigate } from "react-router-dom";

const RedirectedRoute = ({ children }) => {
  const token = localStorage.getItem("token");

  // Jika token ada, arahkan ke halaman utama
  return token ? <Navigate to="/home" /> : children;
};

export default RedirectedRoute;
