
import React from 'react';
import styled from 'styled-components';

const ChatInput = ({input,setInputHandler,handleSend}) => {
    return (
        <Input
            value ={input}
            onChange={e => setInputHandler(e.target.value)}
            onKeyDown={e => e.key === 'Enter' && handleSend()}
        />
    )
}


const Input = styled.input`
  flex: 1;
  padding: 10px;
  border-radius: 4px;
  border: 1px solid #ccc;
`;

export default ChatInput
