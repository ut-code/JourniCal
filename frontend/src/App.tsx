import { RouterProvider, createBrowserRouter } from "react-router-dom";
import Home from "./pages/Home";
import Page1 from "./pages/page1";
import Page2 from "./pages/page2";
import InfiniteScrollTest from "./pages/infinite_scroll_test";
import { CssBaseline } from "@mui/material";

function App() {
  const router = createBrowserRouter([
    {
      path: "/",
      element: <Home />,
    },
    {
      path: "/page1",
      element: <Page1 />,
    },
    {
      path: "/page2",
      element: <Page2 />,
    },
    {
      path: "/infinite_scroll_test",
      element: <InfiniteScrollTest />,
    },
  ]);

  return (
    <>
      <CssBaseline />
      <RouterProvider router={router} />
    </>
  );
}

export default App;
