import { useEffect, useRef, useState } from "react";

const Chat = () => {
    const [messages, setMessages] = useState([]);
    const [input, setInput] = useState("");
    const ws = useRef(null);
    const user_id = 1; // Ganti sesuai user
    const admin_id = 4; // Ganti sesuai admin
    const isConnected = useRef(false); // Menandai apakah koneksi WebSocket sudah terbuka

    useEffect(() => {
        if (!isConnected.current) {
            console.log("Membuka WebSocket...");
            ws.current = new WebSocket(`ws://localhost:3000/ws/chat?user_id=${user_id}&admin_id=${admin_id}`);

            ws.current.onopen = () => {
                console.log("WebSocket Connected!");
                const welcomeMessage = { user_id, admin_id, message: "Hello from client!" };
                console.log("Mengirim pesan pertama: ", welcomeMessage);
                ws.current.send(JSON.stringify(welcomeMessage));
                isConnected.current = true;
            };

            ws.current.onmessage = (event) => {
                try {
                    const receivedMessage = JSON.parse(event.data);
                    console.log("Pesan diterima dari server: ", receivedMessage);

                    // Hanya menambahkan pesan yang belum ada di daftar messages
                    setMessages((prev) => {
                        // Menghindari duplikasi pesan
                        if (!prev.some(msg => msg.message === receivedMessage.message)) {
                            return [...prev, receivedMessage];
                        }
                        return prev;
                    });
                } catch (error) {
                    console.error("Error parsing message:", error);
                }
            };

            ws.current.onerror = (error) => {
                console.error("WebSocket Error:", error);
            };

            return () => {
                console.log("Menutup WebSocket...");
                if (ws.current) {
                    ws.current.close();
                }
            };
        }
    }, []); // Hanya dijalankan sekali saat komponen mount

    const sendMessage = () => {
        if (ws.current && input.trim()) {
            const messageData = { user_id, admin_id, message: input };
            console.log("Mengirim pesan dari input: ", messageData);
            ws.current.send(JSON.stringify(messageData));

            // Menambahkan pesan ke state untuk ditampilkan
            setMessages((prev) => [...prev, messageData]);

            setInput(""); // Mengosongkan input setelah pesan dikirim
        }
    };

    return (
        <div>
            <h2>Chat</h2>
            <div>
                {messages.map((msg, index) => (
                    <p key={index}><strong>{msg.user_id}:</strong> {msg.message}</p>
                ))}
            </div>
            <input
                type="text"
                value={input}
                onChange={(e) => setInput(e.target.value)}
                placeholder="Type a message..."
            />
            <button onClick={sendMessage}>Send</button>
        </div>
    );
};

export default Chat;
