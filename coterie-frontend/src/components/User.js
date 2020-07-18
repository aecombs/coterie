import React from 'react';
import PropTypes from 'prop-types';

const User = (props) => {
  return (
    <div className="w-100 ">
      
      <img src={props.avatar} alt="Avatar" className="w-25" />
      <h5 className="">{props.name}</h5>
      <p className="open-sans">{props.bio ? props.bio : "You don't have a bio yet!" }</p>
    </div>
  )
}

User.propTypes = {
  id: PropTypes.number,
  name: PropTypes.string,
  email: PropTypes.string,
  bio: PropTypes.string,
  avatar: PropTypes.string
};

export default User;