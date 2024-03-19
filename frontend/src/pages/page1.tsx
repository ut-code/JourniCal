import React from "react";
import { Link } from "react-router-dom";

const Page1: React.FC = () => {
  return (
    <>
      <ul>
        <li>
          <Link to="/">Home</Link>
        </li>
        <li>Page1 (Here)</li>
        <li>
          <Link to="/page2">Page2</Link>
        </li>
      </ul>
    </>
  );
};

export default Page1;
