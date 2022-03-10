import React from 'react';
import s from "./Header.module.scss";

export default function Header() {
	return (
		<header className="Header u-contained">
			<img src="/flower.png" id="flower-logo" />
			<div className={s['logo']}>Milton</div>
		</header>
	);
}
