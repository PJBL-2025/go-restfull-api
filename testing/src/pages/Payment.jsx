import React, { useState, useEffect } from "react";
import axios from "axios";

const Payment = () => {
    const [paymentStatus, setPaymentStatus] = useState(null);

    useEffect(() => {
        // Memuat snap.js dari Midtrans setelah komponen di-mount
        const script = document.createElement("script");
        script.src = "https://app.sandbox.midtrans.com/snap/snap.js";
        script.type = "text/javascript";
        script.async = true;
        script.onload = () => console.log("Snap.js loaded successfully.");
        script.onerror = () => console.error("Failed to load Snap.js");
        document.body.appendChild(script);

        return () => {
            document.body.removeChild(script);
        };
    }, []);

    const   handleUpdateStatus = async (id, status) => {
        try{
            const response = await axios.patch(`http://localhost:3000/api/payment/${id}`,{status: status} ,{
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Impyb2NrZXRAZXhhbXBsZS5jb20iLCJleHAiOjE3NDM1NTg2MTYsImlkIjoxLCJyb2xlIjoidXNlciJ9.BiITaB7urkChyqTXd2tW-EsJ_zRPbOVPNsGHtvL7Tt8"
                }
            })
        } catch (e){
            console.log(e)
        }
    }

    const handlePayment = async () => {
        await axios.post(
            "http://localhost:3000/api/payment",
            {
                total_price: 100000,
                product_id: 1,
                quantity: 2,
                address_id: 1,
                payment_method: "credit_card"
            },
            {
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Impyb2NrZXRAZXhhbXBsZS5jb20iLCJleHAiOjE3NDM1NTg2MTYsImlkIjoxLCJyb2xlIjoidXNlciJ9.BiITaB7urkChyqTXd2tW-EsJ_zRPbOVPNsGHtvL7Tt8"
                }
            }
        )
            .then((response) => {
                const snapToken = response.data.data.snap_token;
                const updateStatus = response.data.data.order.id
                console.log(response.data.data.order.id)
                if (snapToken) {
                    if (window.snap) {
                        window.snap.pay(snapToken, {
                            onSuccess: function (result) {
                                handleUpdateStatus(updateStatus, "success")
                                console.log(result);
                                setPaymentStatus("Berhasil");
                            },
                            onPending: function (result) {
                                console.log(result);
                                handleUpdateStatus(updateStatus, "pending")
                                setPaymentStatus("Menunggu Pembayaran");
                            },
                            onError: function (result) {
                                handleUpdateStatus(updateStatus, "failed")
                                console.log(result);
                                setPaymentStatus("Gagal");
                            }
                        });
                    } else {
                        console.error("Snap.js is not loaded.");
                        alert("Snap.js not loaded properly.");
                    }
                } else {
                    alert("Gagal mendapatkan token pembayaran!");
                }
            })
            .catch((error) => {
                console.error("Error:", error);
                alert("Terjadi kesalahan saat membuat pembayaran.");
                setPaymentStatus("Error");
            });
    };

    return (
        <div>
            <h2>Pembayaran Midtrans</h2>
            <p>Total Harga: <strong>Rp 100.000</strong></p>

            <button id="pay-button" onClick={handlePayment}>Bayar Sekarang</button>

            {paymentStatus && (
                <div>
                    <h3>Status Pembayaran: {paymentStatus}</h3>
                </div>
            )}
        </div>
    );
};

export default Payment;
