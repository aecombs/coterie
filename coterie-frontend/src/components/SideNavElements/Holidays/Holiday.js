import React, { useState } from 'react';
import PropTypes from 'prop-types';
import HolidayForm from './HolidayForm';

const Holiday = (props) => {
  const [updateMode, setUpdateMode] = useState(false);

  const formatDate = (dateStr) => {
    const formattedDate = new Date(dateStr).toUTCString();
    return formattedDate.substring(0,17)
  }

  const updateHoliday = (holObj) => {
    props.submitHolidayCallback(holObj);
  }

  const deleteHoliday = () => {
    props.deleteHolidayCallback(props.id);
  }

  const setUpdate = () => {
    setUpdateMode(!updateMode);
  }


  return (
    <section>
      <button onClick={setUpdate} className={ updateMode ? "hidden" : "btn list-group-item list-group-item-action"}>
        <div className="card-body row justify-content-between">
          <div>
            <h5 className="card-title font-weight-bolder text-left">{props.name}</h5>
            <p className="open-sans card-text text-left">{props.description}</p>
          </div>
          <div>
            <p className="open-sans card-text text-right">{formatDate(props.date)}</p>
          </div>
        </div>
      </button>
      <HolidayForm 
        id={props.id}
        name={props.name}
        date={props.date}
        description={props.description}
        orgID={props.orgID}
        visibility={updateMode}
        submitHolidayCallback={updateHoliday}
        onSubmitCallback={setUpdate}
        deleteHolidayCallback={deleteHoliday}
        />
    </section>
  )
}

Holiday.propTypes = {
  id: PropTypes.number,
  name: PropTypes.string,
  date: PropTypes.string,
  description: PropTypes.string,
  orgID: PropTypes.number
};

export default Holiday;