
import Navigation from './Nav';
import { render } from '@testing-library/react'
import userEvent from '@testing-library/user-event';

let el;
beforeEach(() => {
    el = render(<Navigation />)
})

describe("DOM tests", () => {
    it('should be in the document', () => {
        const { getByTestId } = el
        const parentDiv = getByTestId('navigation')
        expect(parentDiv).toBeInTheDocument();
    });
    it('should have a logo that links to the homepage', () => {
        const { getByText } = el
        const logo = getByText('deskmate');
        expect(logo).toBeInTheDocument();
        expect(logo).toHaveAttribute('href', '/');
    });
    it('should have a slackconnect and zendesk connect button', () => {
        const { getByTestId } = el
        const slack = getByTestId('slack_connect');
        const zen = getByTestId('zendesk_connect');
        expect(slack).toBeInTheDocument();
        expect(zen).toBeInTheDocument();
    });
});