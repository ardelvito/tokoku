import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Login from "./pages/Login";
import Register from "./pages/Register";
import Home from "./pages/Home";
import Transactions from "./pages/Transactions";
import Cart from "./pages/Cart";
import Profile from "./pages/Profile/Profile";
import EditProfile from "./pages/Profile/EditProfile";
import DashboardStore  from "./pages/Profile/DashboardStore";
import CreateStore  from "./pages/Profile/CreateStore";
import CreateProduct from './pages/Profile/CreateProduct';
import ProtectedRoute from "./ProtectedRoute";
import RedirectedRoute from "./RedirectedRoute";



function App() {
  return (
  <Router>
      <Routes>
        {/* Protected Route untuk halaman yang memerlukan autentikasi */}
        <Route
          path="/home"
          element={
            <ProtectedRoute>
              <Home />
            </ProtectedRoute>
          }
        />
        <Route
          path="/transactions"
          element={
            <ProtectedRoute>
              <Transactions />
            </ProtectedRoute>
          }
        />
        <Route
          path="/cart"
          element={
            <ProtectedRoute>
              <Cart />
            </ProtectedRoute>
          }
        />
        <Route
          path="/profile/:id"
          element={
            <ProtectedRoute>
              <Profile />
            </ProtectedRoute>
          }
        />
        <Route
          path="/editprofile"
          element={
            <ProtectedRoute>
              <EditProfile />
            </ProtectedRoute>
          }
        />
        <Route
          path="/myproducts/:id"
          element={
            <ProtectedRoute>
              <DashboardStore  />
            </ProtectedRoute>
          }
        />

        <Route
          path="/createproduct"
          element={
            <ProtectedRoute>
              <CreateProduct  />
            </ProtectedRoute>
          }
        />
        <Route
          path="/createstore"
          element={
            <ProtectedRoute>
              <CreateStore  />
            </ProtectedRoute>
          }
        />
        
        {/* Route untuk Login dan Register tanpa perlindungan */}
        <Route
          path="/login"
          element={
            <RedirectedRoute>
              <Login />
            </RedirectedRoute>
          }
        />
        <Route
          path="/register"
          element={
            <RedirectedRoute>
              <Register />
            </RedirectedRoute>
          }
        />
      </Routes>
    </Router>
  );
}

export default App;
