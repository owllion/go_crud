import React, { useState, useEffect, useRef } from 'react';
import styled from 'styled-components';
import SendButton from '../components/sendBtn';
import ChatInput from '../components/chatInput';
import backgroundImage from '../assets/2.jpg';
import ava1 from '../assets/8.jpg'
import ava2 from '../assets/9.jpg'


const Chatroom = () => {
  const [messages, setMessages] = useState([]);
  const [input, setInput] = useState('');
  //TODO: 新增websocket
  const ws = useRef(null);

  useEffect(() => {
    // 初始化 WebSocket 連接
    ws.current = new WebSocket('ws://0.0.0.0:8080/chat/');

    // 監聽接收到的消息
    ws.current.onmessage = (event) => {
      const message = event.data;
      setMessages((prevMessages) => [...prevMessages, message]);
    };

    // 清理操作: 斷開 WebSocket 連接
    return () => {
      if (ws.current) { //NOTE: 啥時關閉???
        ws.current.close();
      }
    };
  }, []);

  const handleSend = () => {
    if (input.trim() !== '') {
      setMessages([...messages, input]);
      setInput('');
    }
  };
  const setInputHandler = (msg) => setInput(msg)
  const isMyMsg = (msg) => msg.length > 5
  

  return (
    <ChatContainer>
      <MessagesContainer>
      {messages.map((message, index) => (
        <MessageWrapper key={index} isOwnMessage={isMyMsg(message)}>
            <MessageBox isOwnMessage={isMyMsg(message)}>
                <img src={isMyMsg(message) ? ava1 : ava2} alt="avatar" />
                {message}
            </MessageBox>
        </MessageWrapper>
      ))}
      </MessagesContainer>
      <InputContainer>
        <ChatInput
            input ={input}
            handleSend = {handleSend}
            setInputHandler = {setInputHandler}
        />
        <SendButton handleSend={handleSend}>Send</SendButton>
      </InputContainer>
        <ClearBtnContainer>
            <ClearBtn onClick={() => setMessages([])}>Clear Room</ClearBtn>
        </ClearBtnContainer>
    </ChatContainer>
  );
};

const ClearBtnContainer = styled.div`
    width: 100%;
    padding-top: 1.5rem;
`

const ClearBtn = styled.button`
    padding: 10px 15px;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    background-color: #000000;
    color: white;
    width: 100%;
`
const ChatContainer = styled.div`
  display: flex;
  flex-direction: column;
  width: 400px;
  height: 500px;
  border-radius: 8px;
  padding: 20px;
  overflow-y: auto;
  background:  url(${backgroundImage});
  background-size: cover; 
  background-position: center; 
  background-color: rgba(0, 0, 0, 0.5); 
`;

const MessagesContainer = styled.div`
  flex: 1;
  overflow-y: auto;
`;

const MessageWrapper = styled.div`
  display: flex;
  align-items: center;  
  margin-bottom: 8px;
  align-self: ${props => (props.isOwnMessage ? 'flex-end' : 'flex-start')};
`;

const MessageBox = styled.div`
  display: flex;
  align-items: center;
  padding: 5px 10px;
  border-radius: 4px;
  background-color: ${props => (props.isOwnMessage ? '#e6f7ff' : '#f5f5f5')};

  img {
    border-radius: 50%;
    width: 30px;
    height: 30px;
    margin-right: 10px;
    object-fit: cover;
  }
`;


const InputContainer = styled.div`
  display: flex;
  margin-top: 10px;
`;


export default Chatroom;
