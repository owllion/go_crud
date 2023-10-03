import { lazy } from "react";
import { useRoutes } from "react-router-dom";
import ProtectedRoute from "../components/Route/ProtectedRoute";

const Home = lazy(() => import("../pages/Home"));

export const RouteConfig = () => {
  let element = useRoutes([
    {
      path: "/",
      element: <Home />,
    },
  ]);

  return element;
};