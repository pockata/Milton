import React from 'react';
import Header from '../components/Header';
import Footer from '../components/Footer';

/**
 * @param {object} props
 * @param {React.ReactNode} props.children
 **/
export default function Layout(props) {
	return (
		<div className="container">
			<Header />
			<main className="u-contained">{props.children}</main>
			<Footer />
		</div>
	);
}

