import React from "react";
import { Link } from "react-router-dom";

const Home: React.FC = () => {
  return (
    <>
      <ul>
        <li>Home (Here)</li>
        <li>
          <Link to="/calendar">Calendar</Link>
        </li>
        <li>
          <Link to="/journal">Journal</Link>
        </li>
      </ul>
    </>
  );
};

export default Home;
