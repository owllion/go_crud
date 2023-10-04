import React, { useState, useEffect } from "react";
import "./styles.css";
import Chatroom from "./pages/chatroom";
import styled from 'styled-components';


export default function App() {
  const [navSize, setnavSize] = useState("10rem");
  const [navColor, setnavColor] = useState("transparent");
  const listenScrollEvent = () => {
    window.scrollY > 10 ? setnavColor("#252734") : setnavColor("transparent");
    window.scrollY > 10 ? setnavSize("5rem") : setnavSize("10rem");
  };
  useEffect(() => {
    window.addEventListener("scroll", listenScrollEvent);
  });

  return (
    <AppContainer>
      {/* <nav
        style={{
          backgroundColor: navColor,
          height: navSize,
          transition: "all 1s"
        }}
      >
        <ul>
          <li>Home</li>
          <li>About</li>
          <li>Project</li>
          <li>Skills</li>
          <li>Contact </li>
        </ul>
      </nav> */}

      <Chatroom/>
    </AppContainer>
  );
}


const AppContainer = styled.div`
  display:flex;
  align-itmes: center;
  justify-content: center;
  padding: 3rem;


`