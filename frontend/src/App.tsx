import { RouterProvider, createBrowserRouter } from "react-router-dom";
import Home from "./pages/Home";
import Page1 from "./pages/page1";
import Page2 from "./pages/page2";

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
  ]);

  return (
    <>
      <RouterProvider router={router} />
    </>
  );
}

export default App;
