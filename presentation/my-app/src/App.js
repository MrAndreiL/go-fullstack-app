import React, { useState } from "react";
import './App.css';
import  { Login } from "./login"
import  { Register } from "./register"
import { Main } from "./main_page"

function App() {
  const [currentPage, setCurrentForm] = useState('login');

  const toggleForm = (formName) => {
    setCurrentForm(formName);
  }

  return (
    <div className="App">
      {
          (() => {
                if (currentPage === 'login') {
                    return (
                        <Login onFormSwitch={toggleForm} />
                    )
                } else if (currentPage === 'register') {
                    return (
                        <Register onFormSwitch={toggleForm} />
                    )
                } else {
                    return (
                        <Main />
                    )
                }
          })()
      }
    </div>
  );
}

export default App;
