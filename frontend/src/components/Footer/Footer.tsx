import React from "react";
import './styles.css';

export const Footer = () => {
    return (
        <footer className="footer">
            <div className="footer__links">
                <div className="footer__links_container">
                    <p className="footer__links_title">Useful links</p>
                    <div className="footer__links_subcontainer">
                        <a className="footer__link" href="#">Tracks Table</a>
                        <a className="footer__link" href="#">Actual Sites Table</a>
                        <a className="footer__link" href="#">Hostinger</a>
                        <a className="footer__link" href="#">NameCheap</a>
                    </div>
                </div>
                <div className="footer__links_container">
                    <p className="footer__links_title">Tracking Sites</p>
                    <div className="footer__links_subcontainer">
                        <a className="footer__link" href="#">metaldeliv</a>
                        <a className="footer__link" href="#">metaldeliv</a>
                        <a className="footer__link" href="#">metaldeliv</a>
                        <a className="footer__link" href="#">metaldeliv</a>
                    </div>
                </div>
                <div className="footer__links_container">
                    <p className="footer__links_title">Links to forums</p>
                    <div className="footer__links_subcontainer">
                        <a className="footer__link" href="#">link #1</a>
                        <a className="footer__link" href="#">link #2</a>
                        <a className="footer__link" href="#">link #3</a>
                        <a className="footer__link" href="#">link #4</a>
                    </div>
                </div>
                <div className="footer__links_container">
                    <p className="footer__links_title">Manuals</p>
                    <div className="footer__links_subcontainer">
                        <a className="footer__link" href="#">link #1</a>
                        <a className="footer__link" href="#">link #2</a>
                        <a className="footer__link" href="#">link #3</a>
                        <a className="footer__link" href="#">link #4</a>
                    </div>
                </div>
                <div className="footer__links_container">
                    <p className="footer__links_title">Contacts</p>
                    <div className="footer__links_subcontainer-contacts">
                        <a className="footer__link" href="#">Jabber: 123123</a>
                        <a className="footer__link" href="#">Telegramm: 12312321</a>
                    </div>
                </div>
            </div>
        </footer>
    )
}