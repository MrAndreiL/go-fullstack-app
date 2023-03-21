import React, { useState } from "react";

export const Main = (props) => {
    const [data, setData] = useState('');
    const [message, setMessage] = useState('');
    
    const handleSubmit = async (e) => {
        e.preventDefault();
        try { 
            let res = await fetch("http://localhost:8083/apis/plagiarism", {
                method: "POST",
                body: JSON.stringify({
                    data: data,
                }),
            });
            let resJson = await res.json();
            setMessage(resJson.plagPercent.toString())
            console.log(resJson)
        } catch (error) {
            console.log(error);
        }
    }

    return (
        <div className="data-form">
            <form className="login-form" onSubmit={handleSubmit}>
                 <textarea value={data} onChange={(e) => setData(e.target.value)} type="data" placeholder="Place your text here!" id="text" name="data"/>
                 <button type="submit">Submit</button>
                 <div className="message">{message ? <p>{message}% plagiarized</p> : null}</div>
            </form>
        </div>
    )
}
