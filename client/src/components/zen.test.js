import { ZendeskConnect } from './ZendeskConnect'
import { render, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import axios from 'axios'

let el;
jest.mock('axios');


describe("on render", () => {
    beforeEach(() => {
        axios.get.mockResolvedValue({ data: true });
    });
    it('should call axios', async () => {
        el = render(<ZendeskConnect />)
        await waitFor(() => expect(axios.get).toHaveBeenCalled());
    });
    it('should have a state of "not connected"', () => {
        axios.get.mockResolvedValue({ data: false });
        el = render(<ZendeskConnect />)
        const { queryByText } = el;
        const label = queryByText('Not Connected: Zendesk');
        expect(label).not.toBe(null);
    });
    it('should update the label on connecting to slack', async () => {
        el = render(<ZendeskConnect />)
        const { queryByText } = el;
        const label = queryByText('Connected: Zendesk');
        expect(axios.get).toHaveBeenCalled();
        () => expect(label).not.toBe(null);
    })
});

describe('on click', () => {
    beforeEach(() => {
        axios.get.mockResolvedValue({ data: false });
        axios.post.mockResolvedValue({});
    });
    it('should POST on click', async () => {
        const { queryByRole } = render(<ZendeskConnect />)
        const btn = queryByRole('button');
        userEvent.click(btn);
        await waitFor(() => expect(axios.post).toHaveBeenCalled());
    })
})