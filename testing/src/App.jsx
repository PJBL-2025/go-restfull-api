import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import {createBrowserRouter, RouterProvider} from "react-router-dom";
import Chat from "./pages/Chat.jsx";
import Payment from "./pages/Payment.jsx";


function App() {
    const router = createBrowserRouter([
        {
            path:"/chat",
            element: <Chat/>
        },
        {
            path:"/payment",
            element:<Payment/>
        }
    ])
  return (
    <>
        <RouterProvider router={router}/>
    </>
  )
}

export default App
