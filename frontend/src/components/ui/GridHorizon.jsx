import { useEffect, useRef } from 'react';
import './GridHorizon.css';

export default function GridHorizon() {
  return (
    <div className="grid-horizon">
      <div className="sky-background"></div>
      <div className="stars-field">
        <div className="stars-layer"></div>
      </div>
      <div className="grid-floor"></div>
      <div className="grid-vertical-lines"></div>
      <div className="horizon-line"></div>
    </div>
  );
} 