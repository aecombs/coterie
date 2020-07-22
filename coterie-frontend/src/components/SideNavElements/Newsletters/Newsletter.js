import React, { useState } from 'react';
import PropTypes from 'prop-types';
import NewsletterForm from './NewsletterForm';

const Newsletter = (props) => {
  const [updateMode, setUpdateMode] = useState(false);

  const formatDate = (dateStr) => {
    const formattedDate = new Date(dateStr).toUTCString();
    return formattedDate.substring(0,17)
  }

  const updateNewsletter = (newsObj) => {
    props.submitNewsletterCallback(newsObj)
  }

  const deleteNewsletter = () => {
    props.deleteNewsletterCallback(props.id);
  }

  const setUpdate = () => {
    setUpdateMode(!updateMode);
  }

  return (
    <section>
      <button onClick={setUpdate} className={ updateMode ? "hidden" : "btn list-group-item list-group-item-action"}>
        <div className="card-body row justify-content-between">
          <div>
            <h5 className="card-title font-weight-bolder text-left">{props.header}</h5>
            <p className="open-sans card-text text-left">{props.description}</p>
          </div>
          <div>
            <p className="open-sans card-text text-right">{formatDate(props.date)}</p>
          </div>
        </div>
      </button>
      <NewsletterForm 
        id={props.id}
        header={props.header}
        description={props.description}
        date={props.date}
        orgID={props.orgID}
        visibility={updateMode}
        submitNewsletterCallback={updateNewsletter}
        onSubmitCallback={setUpdate}
        deleteNewsletterCallback={deleteNewsletter}
        />
    </section>
  )
}

Newsletter.propTypes = {
  id: PropTypes.number,
  header: PropTypes.string,
  description: PropTypes.string,
  date: PropTypes.string,
  orgID: PropTypes.number
};

export default Newsletter;