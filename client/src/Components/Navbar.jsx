import React from 'react';

const Navbar = ({showSearchBar}) => {
	return (
		<div className="navbar">
			<div className="navbar_left">
				<i class="fas fa-bars"></i>
				CHATEX
			</div>
			{showSearchBar ? (
				<div className="navbar_middle">
					<label htmlFor="search" className="navbar_middle_search">
						<i class="fas fa-search"></i>
						<input type="text" name="search" id="search" placeholder="Search here..." />
					</label>
				</div>
			) : (<></>)}
			<div className="navbar_right">
				{/* settings toggle */}
				<i class="fas fa-cog"></i>
				{/* logout button */}
				<button className="navbar_right_logout">Logout</button>
				<i class="fas fa-sign-out-alt"></i>
			</div>
		</div>
	)
};

export default Navbar;
