import React, { useState } from 'react';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';
import OrganizationContainer from './Dashboard/OrganizationContainer';
import Profile from './SideNavElements/Profile/Profile';
import Newsletters from './SideNavElements/Newsletters/Newsletters';
import Events from './SideNavElements/Events/Events';
import Holidays from './SideNavElements/Holidays/Holidays';
import Scriptures from './SideNavElements/Scriptures/Scriptures';
import OrgForm from './Dashboard/OrgForm';
import axios from 'axios';

const SideNav = () => {
  const [orgID, setOrgID] = useState(null);
  const [errorMessage, setErrorMessage] = useState(null);

  const userID = sessionStorage.getItem('userID');

  const url = `${process.env.REACT_APP_API_BASE_URL}/users/${userID}/organizations/`



  const addOrg = (orgObj) => {
    axios.post(url, orgObj)
    .then((response) => {
      setErrorMessage(`organization ${orgObj["name"]} added`);
      //TODO: Update to use intended redirect URL from response
      window.location.assign(`${process.env.REACT_APP_BASE_URL}/dashboard`);
    })
    
    .catch((error) => {
      setErrorMessage(error.message);
      console.log(`Unable to add org: ${errorMessage}`);
    })
  }

  const updateOrgSession = (orgID) => {
    if (orgID) {
      sessionStorage.setItem('orgID', orgID);
      setOrgID(sessionStorage.getItem('orgID'));
    }
  }




  return (
    <Router>
    <div className="d-flex w-100">
      <div className="side-nav">
        <nav className="side-nav">
          <ul className="nav flex-column">
            <li className="nav-item nav-link text-decoration-none"><Link to={'/dashboard'} className="">Dashboard</Link></li>
            <li className="nav-item nav-link"><Link to={'/dashboard/profile'} className="">Profile</Link></li>
            <li className="nav-item nav-link"><Link to={'/dashboard/scriptures'} className="">Scriptures</Link></li>
            <li className="nav-item nav-link"><Link to={'/dashboard/holidays'} className="">Holidays</Link></li>
            <li className="nav-item nav-link"><Link to={'/dashboard/events'} className="">Events</Link></li>
            <li className="nav-item nav-link"><Link to={'/dashboard/newsletters'} className="">Newsletters</Link></li>
          </ul>
        </nav>
      </div>
      <div className="dashboard-display">
        <Switch>
        {/* <Route exact path='/dashboard' component={Organization} /> */}
        <Route 
            exact path='/dashboard' 
            render={(props) => (
              <OrganizationContainer {...props}
              userID={userID}
              addOrgCallback={addOrg}
              tellParentID={updateOrgSession}
              />
            )}
          />

          <Route exact path='/dashboard/profile'
           render={(props) => (
            <Profile {...props} 
            userID={userID}
            />
          )} />
          <Route 
            exact path='/dashboard/newsletters' 
            render={(props) => (
              <Newsletters {...props} 
              userID={userID}
              orgID={orgID}
              />
            )}
          />
          <Route 
            exact path='/dashboard/events' 
            render={(props) => (
              <Events {...props} 
              userID={userID}
              orgID={orgID}
              />
            )}
          />
          
          <Route 
            exact path='/dashboard/holidays' 
            render={(props) => (
              <Holidays {...props} 
              userID={userID}
              orgID={orgID}
              />
            )}
          />
          <Route 
            exact path='/dashboard/scriptures' 
            render={(props) => (
              <Scriptures {...props} 
              userID={userID}
              orgID={orgID}
              />
            )}
          />
          

          <Route 
            exact path='/dashboard/new'
            render={(props) => (
              <OrgForm {...props} 
              addOrgCallback={addOrg}
              />
            )}
          />

        </Switch>
      </div>
    </div>
  </Router>
  )
}

export default SideNav;