import React, {  useState  } from 'react';
import { 
    Collapse,
    Navbar,
    NavbarToggler,
    NavbarBrand,
    Nav,
    NavItem,
    NavLink,
    UncontrolledDropdown,
    DropdownToggle,
    DropdownMenu,
    DropdownItem,
    NavbarText
  } from 'reactstrap';
import {SlackConnect} from './SlackConnect.js';
import { ZendeskConnect } from './ZendeskConnect.js';

const Navigation = () => {
  const [isOpen, setIsOpen] = useState(false);

  const toggle = () => setIsOpen(!isOpen);

  return (
    <div data-testid="navigation">
      <Navbar color="light" light expand="md">
        <NavbarBrand data-testid="logo" href="/">deskmate</NavbarBrand>
        <NavbarToggler data-testid="toggler" onClick={toggle} />
        <Collapse data-testid="collapse" isOpen={isOpen} navbar>
          <NavbarText>
            <SlackConnect />
            <ZendeskConnect />
          </NavbarText>
        </Collapse>
      </Navbar>
    </div>
  );
}

export default Navigation;