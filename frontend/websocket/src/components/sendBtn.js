
import React from 'react';
import styled from 'styled-components';
import { FaTelegramPlane } from "react-icons/fa";

const SendButton = ({handleSend}) => {
    return (
        <Button onClick={handleSend}>
          <FaTelegramPlane/>
        </Button>
    )
}

const Button = styled.button`
  padding: 10px 15px;
  margin-left: 10px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  background-color: #000000;
  color: white;
`;
    
export default SendButton