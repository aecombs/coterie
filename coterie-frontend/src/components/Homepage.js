import React from 'react';

const Homepage = () => {
  return (
    <div className="homepage w-100">
      
      <div className="container homepage-container">
        <h5 className="my-5 text-center">Welcome!</h5>
        <hr/>
        <div className="row my-5">
          <div className="col text-right">
          <p className="open-sans">Do you identify with this flamingo here on the right?</p>
          <p className="open-sans">Unique, important, uniqely important. A touch contemplative. Perhaps you're the chosen leader of the True people. Maybe you've even heard the Truth from on high -- a Truth that not everyone is willing to accept. You might be working endlessly to spread this word and bring those that accept this Truth to salvation. You might've thought about giving up this endeavor that the universe has compelled you to complete.</p>
          <p className="open-sans">We know it's not easy being in your position. From performing rituals and blessings to hesitantly punishing those in your flock who doubt the Truth, we know you do your best to make them fit for the next world.</p>
          <p className="open-sans">We understand that you need help in managing this organization of yours. After all, teaching your flock the ways of righteousness is not at easy task!</p>
          <h4 className="open-sans text-center mt-5 mb-4">That's what we're here for.</h4>
          </div>
          
          <div className="col text-right">
            <img className="homepage-small-img" alt="Flamingo in Focus" src="https://images.unsplash.com/photo-1595036280307-faeb2808724a?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=2550&q=80"/>
          </div>
        </div>
        <hr/>
        <h3 className="open-sans text-center my-5">Our mission is to help you help others.</h3>
        
        <div className="row mb-5">
          <div className="col">
            <img className="homepage-small-img" alt="Flamingo interacting with another flamingo" src="https://images.unsplash.com/photo-1573492986146-b8a59cd771d0?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjQ0MTA5fQ&auto=format&fit=crop&w=2550&q=80"/>
          </div>
          <div className="col">
            <p className="open-sans">We know how important it is that everyone on this Earth hears and understands your message. We get it! It's quite literally a matter of life or death.</p>
            <p className="open-sans bold">In order to make this difficult task easier for you, we developed <strong className="open-sans">Coterie</strong>, an unique, online tool to help you organize your organization.</p>
            <p className="open-sans">Coterie makes it possible for you to document your Truth, keep track of your organization's unique holidays, and so much more.</p>
            <h6 className="open-sans font-weight-bolder">On a budget but don't trust the banks to track your money?</h6>
            <p className="open-sans">With Coterie, you can upload the funds raised by individual members to dynamically calculate how much money you have.</p>
            <h6 className="open-sans font-weight-bolder">Only got one written copy of your scriptures?</h6>
            <p className="open-sans">You can document individual scriptures and their corresponding chapters on our platform. You'll never have to worry about blood splatters from your most recent sacrifical ritual obscuring your written Truth when you can reference the digital copy!</p>
          </div>
        </div>
        <hr/>
        {/* <h4 className="mt-3 text-center">Coterie</h4>
        <p className="open-sans mb-3 text-center">The online tool to help you organize your organization.</p> */}
      </div>
      <div className="container my-5">
        <img className="w-100 homepage-large-img" alt="A Flamboyance of Flamingos" src="https://images.unsplash.com/photo-1515048397571-d8d6729fb49c?ixlib=rb-1.2.1&auto=format&fit=crop&w=2689&q=80"/>
      </div>
      <div className="container mb-5">
        <div className="my-5">
          <h4 className="text-center font-weight-bold mb-4">Why Flamingos?</h4>
          <p className="open-sans">Much like humans, flamingos are very social creatures. They tend to live in large colonies, with numbers ranging from 50 up to the thousands. They often communicate using non-verbal cues and engage in group activities. Flamingos form strong pair bonds with others of either the same sex or opposite sex, and they learn to recognize the calls of their parents and of their children so that they can hear them even in large colonies. As for the young, they form groups known as crèches, and the entire colony will take on the burden of caring for them.</p>
        </div>
          
        <div className="row mb-5">
          <div className="col">
            <img className="homepage-small-img" alt="Closeup of Flamingo" src="https://images.unsplash.com/photo-1492201722955-33e840ece203?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=2555&q=80"/>
          </div>
          <div className="col">
            <p className="open-sans">Mind you, we're not here to teach you about flamingos, but it is interesting to point out the similarities. More than anything, this is an exercise in explaining the theme post-selection. The color scheme was really the reason for choosing flamingos.</p>
            <p className="open-sans">So, what <em className="open-sans">are</em> we here to do?</p>
            <p className="open-sans">We're here to help.</p>
            <p className="open-sans">We're here to offer a helping hand during this time of struggle and imminent eternal damnation.</p>
            <p className="open-sans">We're here for <em className="open-sans">you</em>.</p>
            <p className="open-sans">March on, glorious leader.</p>
          </div>
          
        </div>
        <hr/>
      </div>
      
      <div className="container text-center my-5">
        <h4 className=" mt-5">Coterie.</h4>
        <h4 className="my-5">We welcome you.</h4>
        <a href="/login" className="text-decoration-none text-white"><button className="btn btn-secondary w-50">Sign up today.</button></a>
      </div>
    </div>
  )
}

export default Homepage;