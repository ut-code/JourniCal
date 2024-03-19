import React from "react";
import { Link } from "react-router-dom";

const Page2: React.FC = () => {
  return (
    <>
      <ul>
        <li>
          <Link to="/">Home</Link>
        </li>
        <li>
          <Link to="/page1">page1</Link>
        </li>
        <li>Page2 (Here)</li>
      </ul>
    </>
  );
};

export default Page2;
