import React from 'react';
import ChatWindow from '../Components/ChatWindow';
import Navbar from '../Components/Navbar';
import Sidebar from '../Components/Sidebar';

const MainPage = () => {
	return <div className="dashboard">
		<Navbar showSearchBar={true} />
		<div className="dashboard_main">
			<Sidebar />
			<ChatWindow />
		</div>
	</div>;
};

export default MainPage;
