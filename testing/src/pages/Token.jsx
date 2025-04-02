import {useRef} from "react";

const Token = () => {
    const tokenRef = useRef()
    const handleSubmit = () => {
        localStorage.setItem("token", tokenRef.current.value)
    }
    return(
        <>
            <input ref={tokenRef}/>
            <button onClick={handleSubmit}/>
        </>
    )
}

export  default  Token