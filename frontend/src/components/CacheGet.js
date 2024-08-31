import React, { useState } from 'react';
import axios from 'axios';

function CacheGet() {
    const [key, setKey] = useState('');
    const [value, setValue] = useState('');
    const [error, setError] = useState('');

    const handleGet = async () => {
        try {
            const response = await axios.get(`http://localhost:3001/cache/${key}`);
            setValue(response.data.value);
            setError('');
        } catch (err) {
            setError('Key not found or an error occurred');
            setValue('');
        }
    };

    return (
        <div>
            <h2>Get Cache Value</h2>
            <input
                type="text"
                placeholder="Enter key"
                value={key}
                onChange={(e) => setKey(e.target.value)}
            />
            <button onClick={handleGet}>Get Value</button>
            {value && (
                <div>
                    <strong>Value:</strong> {value}
                </div>
            )}
            {error && (
                <div style={{ color: 'red' }}>
                    {error}
                </div>
            )}
        </div>
    );
}

export default CacheGet;
