import React, { useState } from "react";

export const Register = (props) => {
    const [email, setEmail] = useState('');
    const [pass, setPass] = useState('');
    const [name, setName] = useState('');
    const [username, setUsername] = useState('');
    const [message, setMessage] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            let isStudent = true;
            let res = await fetch("http://localhost:8082/apis/check", {
                method: "POST",
                body: JSON.stringify({
                    email: email,
                }),
            });
            let resJson = await res.json();
            if (res.status !== 200) {
                setMessage("Invalid student email");
                isStudent = false;
            }
            if (isStudent) {
                let res2 = await fetch("http://localhost:8081/apis/register", {
                    method: "POST",
                    body: JSON.stringify({
                        name: name,
                        username: username,
                        email: email,
                        password: pass,
                    }),
                });
                let data = await res2.json();
                if (res2.status !== 201) {
                    setMessage("Invalid credentials.");
                } else {
                    props.onFormSwitch('login');
                }
            }
        } catch (err) {
            console.log(err);
        }
    }

    return (
        <div className="auth-form-container">
            <h2>Register</h2>
        <form className="register-form" onSubmit={handleSubmit}>
            <label htmlFor="name">Full name</label>
            <input value={name} name="name" onChange={(e) => setName(e.target.value)} id="name" placeholder="full Name" />
            <label htmlFor="username">Username</label>
            <input value={username} name="username" onChange={(e) => setUsername(e.target.value)} id="username" placeholder="user name" />
            <label htmlFor="email">email</label>
            <input value={email} onChange={(e) => setEmail(e.target.value)}type="email" placeholder="youremail@gmail.com" id="email" name="email" />
            <label htmlFor="password">password</label>
            <input value={pass} onChange={(e) => setPass(e.target.value)} type="password" placeholder="********" id="password" name="password" />
            <button type="submit">Register</button>
        </form>
        <button className="link-btn" onClick={() => props.onFormSwitch('login')}>Already have an account? Login here.</button>
        <div className="message">{message ? <p>{message}</p> : null}</div>
    </div>
    )
}
