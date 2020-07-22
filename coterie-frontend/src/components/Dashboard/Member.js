import React, { useState } from 'react';
import PropTypes from 'prop-types';
import MemberForm from './MemberForm';

const Member = (props) => {
  const [updateMode, setUpdateMode] = useState(false);


  const formatCurrency = (int) => {
    let modInt = 0;
    if (int !== 0) {
      modInt = int/100
    }
    if (int === undefined || int === null) {
      modInt = 0;
    }
    const formattedInt = new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD' }).format(modInt)
    return formattedInt
  }

  const formatDate = (date) => {
    const month = date.substring(5, 7)
    const day = date.substring(8, 10)
    const year = date.substring(0, 4)
    const string = `${year}-${month}-${day}`
    const formattedDate = new Date(string).toUTCString();
    return formattedDate.substring(0, 17)
  }

  const updateMember = (memObj) => {
    props.submitMemberCallback(memObj)
  }

  const deleteMember = () => {
    props.deleteMemberCallback(props.id);
  }

  const setUpdate = () => {
    setUpdateMode(!updateMode);
  }


  return (
    <section>
      <button onClick={setUpdate} className={ updateMode ? "hidden" : "btn border-0 list-group-item list-group-item-action"}>
        <div className="card-body row justify-content-between">
          <div className="text-left">
            <h5 className="card-title font-weight-bolder text-left">{props.name}</h5>
          </div>
          <div className="row container justify-content-around">
            <div className="text-left">
              <small>Class</small>
              <p className="open-sans card-text text-left">{props.class}</p>
            </div>
            <div className="text-left">
              <small>Birthdate</small>
              <p className="open-sans card-text text-right">{formatDate(props.birthdate)}</p>
            </div>
            <div className="text-left">
              <small>Email</small>
              <p className="open-sans card-text text-right">{props.email}</p>
            </div>
            <div className="text-left">
              <small>Total funds raised</small>
              <p className="open-sans card-text text-right">{formatCurrency(props.fundsRaised)}</p>
            </div>
          </div>
        </div>
      </button>
      <MemberForm 
        id={props.id}
        name={props.name}
        birthdate={props.birthdate}
        class={props.class}
        email={props.email}
        fundsRaised={props.fundsRaised}
        orgID={props.orgID}
        visibility={updateMode}
        submitMemberCallback={updateMember}
        onSubmitCallback={setUpdate}
        deleteMemberCallback={deleteMember}
        />
    </section>
  )
}

Member.propTypes = {
  id: PropTypes.number,
  name: PropTypes.string,
  birthdate: PropTypes.string,
  class: PropTypes.string,
  email: PropTypes.string,
  fundsRaised: PropTypes.number,
  orgID: PropTypes.number
};

export default Member;