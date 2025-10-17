import React from "react";
import './styles.css';
import { NavLink } from "react-router-dom";

export const Header = () => {
    return (
        <header className="header">
            <p className="header__title">
                <span className="header__title-span">Admin Panel</span> by Skadi
            </p>
            <div className="header__links">
                <NavLink 
                    to="/tracks" 
                    className={({ isActive }) => 
                        isActive ? "header__links_el active" : "header__links_el"
                    }
                >
                    TRACKS
                </NavLink>
                <NavLink 
                    to="/api" 
                    className={({ isActive }) => 
                        isActive ? "header__links_el active" : "header__links_el"
                    }
                >
                    API
                </NavLink>
                <NavLink 
                    to="/tg-api" 
                    className={({ isActive }) => 
                        isActive ? "header__links_el active" : "header__links_el"
                    }
                >
                    TG API
                </NavLink>
            </div>
        </header>
    );
};