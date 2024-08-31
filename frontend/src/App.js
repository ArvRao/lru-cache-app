import React from 'react';
import CacheGet from './components/CacheGet';
import CacheSet from './components/CacheSet';
import './App.css';

function App() {
    return (
        <div className="App">
            <header className="App-header">
                <h1>LRU Cache Application</h1>
            </header>
            <main>
                <CacheSet />
                <CacheGet />
            </main>
        </div>
    );
}

export default App;
