import { Navigate } from "react-router-dom";

const ProtectedRoute = ({ children }) => {
  const token = localStorage.getItem("token");

  // Jika tidak ada token, arahkan ke halaman login
  return token ? children : <Navigate to="/login" />;
};

export default ProtectedRoute;
