import React from 'react'

const ContactCard = ({name}) => {
	return (
		<div className="contact_card">
			<div className="contact_card_image"></div>
			<div className="contact_card_name">{name}</div>
		</div>
	)
}

export default ContactCard
