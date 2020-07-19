import React, { useState } from 'react';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';
import OrganizationContainer from './Dashboard/OrganizationContainer';
import Profile from './SideNavElements/Profile/Profile';
import Announcements from './SideNavElements/Announcements/Announcements';
import Events from './SideNavElements/Events/Events';
import Holidays from './SideNavElements/Holidays/Holidays';
import Scriptures from './SideNavElements/Scriptures/Scriptures';
import OrgForm from './Dashboard/OrgForm';
import axios from 'axios';

const SideNav = (props) => {
  const [errorMessage, setErrorMessage] = useState(null);

  const organizationID = props.orgID

  const url = `http://localhost:3000/users/${props.userID}/organizations/`

  const addOrg = (orgObj) => {
    axios.post(url, orgObj)
    .then((response) => {
      setErrorMessage(`organization ${orgObj["name"]} added`);
      //TODO: Update to use intended redirect URL from response
      window.location.assign("http://localhost:3001/dashboard");
    })
    
    .catch((error) => {
      setErrorMessage(error.message);
      console.log(`Unable to add org: ${errorMessage}`);
    })
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
            <li className="nav-item nav-link"><Link to={'/dashboard/announcements'} className="">Announcements</Link></li>
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
              userID={"1"}
              addOrgCallback={addOrg}
              />
            )}
          />

          <Route path='/dashboard/profile'
           render={(props) => (
            <Profile {...props} 
            userID={"1"}
            />
          )} />
          <Route 
            path='/dashboard/announcements' 
            render={(props) => (
              <Announcements {...props} 
              userID={"1"}
              orgID={"1"}
              />
            )}
          />
          <Route 
            path='/dashboard/events' 
            render={(props) => (
              <Events {...props} 
              userID={"1"}
              orgID={"1"}
              />
            )}
          />
          
          <Route 
            path='/dashboard/holidays' 
            render={(props) => (
              <Holidays {...props} 
              userID={"1"}
              orgID={"1"}
              />
            )}
          />
          <Route 
            path='/dashboard/scriptures' 
            render={(props) => (
              <Scriptures {...props} 
              userID={"1"}
              orgID={"1"}
              />
            )}
          />
          

          <Route 
            path='/dashboard/new'
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