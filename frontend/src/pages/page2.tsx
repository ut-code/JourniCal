import React from "react";
import { Link } from "react-router-dom";
import TopBar from "../components/TopBar";

const Page2: React.FC = () => {
  return (
    <>
      <TopBar journalPathName="/page1" calendarPathName="/page2" />
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
