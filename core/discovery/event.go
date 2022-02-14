/*
 * Copyright (C) 2020 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package discovery

// Topic represents the different topics a consumer can subscribe to
const (
	// AppTopicProposalAdded represents newly announced proposal
	AppTopicProposalAdded = "ProposalAdded"
	// AppTopicProposalUpdated represents re-announced proposal
	AppTopicProposalUpdated = "ProposalUpdated"
	// AppTopicProposalRemoved represents newly de-announced proposal
	AppTopicProposalRemoved = "ProposalRemoved"
	// AppTopicProposalAnnounce represent proposal events topic.
	AppTopicProposalAnnounce = "proposalEvent"
)
