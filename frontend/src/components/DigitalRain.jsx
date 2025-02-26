import { useEffect, useRef } from 'react';
import './DigitalRain.css';

export default function DigitalRain() {
  const containerRef = useRef(null);
  const columnsRef = useRef([]);
  
  useEffect(() => {
    if (!containerRef.current) return;
    
    const container = containerRef.current;
    const containerWidth = window.innerWidth;
    const columnCount = Math.floor(containerWidth / 30); // One column every ~30px
    
    // Clear any existing columns
    while (container.firstChild) {
      container.removeChild(container.firstChild);
    }
    
    columnsRef.current = [];
    
    // Create columns
    for (let i = 0; i < columnCount; i++) {
      const column = document.createElement('div');
      column.className = 'rain-column';
      column.style.left = `${i * 30 + Math.random() * 15}px`; // Random offset
      
      // Random characters
      const length = 10 + Math.floor(Math.random() * 15); // Random length
      for (let j = 0; j < length; j++) {
        const char = String.fromCharCode(0x30A0 + Math.random() * 96); // Katakana characters
        const span = document.createElement('span');
        span.textContent = char;
        span.style.opacity = j === 0 ? '1' : '0.8'; // First character brighter
        column.appendChild(span);
        column.appendChild(document.createElement('br'));
      }
      
      // Random animation duration and delay
      const duration = 3 + Math.random() * 8;
      column.style.animationDuration = `${duration}s`;
      column.style.animationDelay = `${Math.random() * 5}s`;
      
      container.appendChild(column);
      columnsRef.current.push(column);
    }
    
    // Handle window resize
    const handleResize = () => {
      // Re-render on significant width changes
      if (Math.abs(window.innerWidth - containerWidth) > 100) {
        // Clear timeout if it exists
        if (columnsRef.current.resizeTimeout) {
          clearTimeout(columnsRef.current.resizeTimeout);
        }
        
        // Set a timeout to avoid too many re-renders
        columnsRef.current.resizeTimeout = setTimeout(() => {
          // Re-run this effect
          columnsRef.current.forEach(col => container.removeChild(col));
          columnsRef.current = [];
          createColumns();
        }, 300);
      }
    };
    
    window.addEventListener('resize', handleResize);
    
    return () => {
      window.removeEventListener('resize', handleResize);
      if (columnsRef.current.resizeTimeout) {
        clearTimeout(columnsRef.current.resizeTimeout);
      }
    };
  }, []);
  
  return <div className="digital-rain-container" ref={containerRef}></div>;
} 