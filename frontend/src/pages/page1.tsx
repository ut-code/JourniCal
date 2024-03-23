import React from "react";
import { Link } from "react-router-dom";
import TopBar from "../components/TopBar";

const Page1: React.FC = () => {
  return (
    <>
      <TopBar journalPathName="/page1" calendarPathName="/page2"/>
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
