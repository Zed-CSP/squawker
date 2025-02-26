import React, { useEffect, useRef } from 'react';
import './CyberpunkLogo.css';

const CyberpunkLogo = () => {
  const logoRef = useRef(null);
  
  useEffect(() => {
    const logo = logoRef.current;
    if (!logo) return;
    
    // Random severe glitch effect
    const randomGlitch = () => {
      const glitchChance = Math.random();
      
      if (glitchChance > 0.85) { // 15% chance of severe glitch
        logo.classList.add('severe-glitch');
        setTimeout(() => {
          logo.classList.remove('severe-glitch');
        }, 200 + Math.random() * 300);
      } else if (glitchChance > 0.6) { // 25% chance of mild glitch
        logo.classList.add('glitching');
        setTimeout(() => {
          logo.classList.remove('glitching');
        }, 100 + Math.random() * 150);
      }
    };
    
    // Set up interval for random glitches
    const glitchInterval = setInterval(randomGlitch, 3000);
    
    // Add hover effects
    const handleMouseEnter = () => {
      clearInterval(glitchInterval);
      logo.classList.add('cyber-hover');
    };
    
    const handleMouseLeave = () => {
      logo.classList.remove('cyber-hover');
      setTimeout(() => {
        // Add one final glitch when leaving
        logo.classList.add('glitching');
        setTimeout(() => {
          logo.classList.remove('glitching');
        }, 200);
      }, 150);
    };
    
    logo.addEventListener('mouseenter', handleMouseEnter);
    logo.addEventListener('mouseleave', handleMouseLeave);
    
    return () => {
      clearInterval(glitchInterval);
      logo.removeEventListener('mouseenter', handleMouseEnter);
      logo.removeEventListener('mouseleave', handleMouseLeave);
    };
  }, []);
  
  return (
    <div className="cyberpunk-logo-wrapper">
      <div className="cyberpunk-logo" ref={logoRef}>
        <div className="logo-inner">
          <span className="logo-text" data-text="S">S</span>
          <span className="logo-text" data-text="Q">Q</span>
          <span className="logo-text" data-text="U">U</span>
          <span className="logo-text" data-text="A">A</span>
          <span className="logo-text" data-text="W">W</span>
          <span className="logo-text" data-text="K">K</span>
          <span className="logo-text" data-text="E">E</span>
          <span className="logo-text" data-text="R">R</span>
        </div>
        <div className="logo-shadow"></div>
      </div>
      <div className="logo-circuit-lines"></div>
      <div className="logo-glow"></div>
    </div>
  );
};

export default CyberpunkLogo; 