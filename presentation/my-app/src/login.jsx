import React, { useState } from "react";

export const Login = (props) => {
    const [username, setUsername] = useState('');
    const [pass, setPass] = useState('');
    const [message, setMessage] = useState('');
    
    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            let res = await fetch("http://localhost:8081/apis/login", {
                method: "POST",
                body: JSON.stringify({
                    username: username,
                    password: pass,
                }),
            });
            let resJson = await res.json();
            if (res.status !== 200) {
                setMessage("Incorrect username or password");
            } else {
                props.onFormSwitch('main');
            }
        } catch (err) {
            console.log(err);
        }
    }

    return (
        <div className="auth-form-container">
            <h2>Login</h2>
            <form className="login-form" onSubmit={handleSubmit}>
                 <label htmlFor="username">username</label>
                 <input value={username} onChange={(e) => setUsername(e.target.value)} type="username" placeholder="myusername" id="username" name="username"/>
                 <label htmlFor="password">password</label>
                 <input value={pass} onChange={(e) => setPass(e.target.value)} type="password" placeholder="****" id="password" name="password"/>
                 <button type="submit">Login</button>
            </form>
            <button className="link-btn" onClick={() => props.onFormSwitch('register')}>Do not have an account? Register here.</button>
            <div className="message">{message ? <p>{message}</p> : null}</div>
        </div>
    )
}
