import React from 'react';
import PropTypes from 'prop-types';

const Form = (props) => {

  const url = `http://localhost:3000/users/${props.userID}/organizations/${props.orgID}/${props.selection}`;

  

  return (
    <div className="">

    </div>
  )
}

Form.propTypes = {
  selection: PropTypes.string,
  userID: PropTypes.number,
  orgID: PropTypes.number
};

export default Form;