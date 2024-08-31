import React, { useState } from 'react';
import axios from 'axios';

function CacheSet() {
    const [key, setKey] = useState('');
    const [value, setValue] = useState('');
    const [ttl, setTtl] = useState('');
    const [message, setMessage] = useState('');

    const handleSet = async () => {
        try {
            await axios.post('http://localhost:3001/cache', {
                key,
                value,
                ttl: ttl ? parseInt(ttl, 10) : undefined,
            });
            setMessage('Key-Value pair set successfully');
        } catch (err) {
            setMessage('An error occurred while setting the key-value pair');
        }
    };

    return (
        <div>
            <h2>Set Cache Value</h2>
            <input
                type="text"
                placeholder="Enter key"
                value={key}
                onChange={(e) => setKey(e.target.value)}
            />
            <input
                type="text"
                placeholder="Enter value"
                value={value}
                onChange={(e) => setValue(e.target.value)}
            />
            <input
                type="text"
                placeholder="Enter TTL (seconds)"
                value={ttl}
                onChange={(e) => setTtl(e.target.value)}
            />
            <button onClick={handleSet}>Set Value</button>
            {message && (
                <div>
                    {message}
                </div>
            )}
        </div>
    );
}

export default CacheSet;
