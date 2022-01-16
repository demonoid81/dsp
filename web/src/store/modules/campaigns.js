import axios from "axios";

const state = {
    campaign: {
        name: '',
        url: '',
        type: [],
        ad: {
            icon: '',
            image: '',
            title: '',
            text: '',
        },
        campaignCountries: [{
            country: "US",
            cpc: 0.001
        }],
        target: {
            os: [],
            browser: []
        }
    },
}

const getters = {
    campaign: state => state.campaign,
    campaignCountries: state => {
        return state.campaign.campaignCountries.map((item, index) => {
            return {
                ...item,
                index: index
            }
        })
    }
}

const mutations = {
    setCampaign(state, campaign) {
        state.campaign = campaign
    },
    AddCampaignCountry(state) {
        state.campaign.campaignCountries.push({
            country: '',
            cpc: 0.0001
        })
    },
    UpdateCampaignCountryItemCountry(state, {country, index}) {
        state.campaign.campaignCountries = state.campaign.campaignCountries.
        map((item, itemIndex) => {
            if (index === itemIndex) {
                return {
                    ...item,
                    country: country
                }
            }
            return {
                ...item
            }
        })
    },
    UpdateCampaignCountryItemCPC(state, {cpc, index}) {
        state.campaign.campaignCountries = state.campaign.campaignCountries.
        map((item, itemIndex) => {
            if (index === itemIndex) {
                return {
                    ...item,
                    cpc: cpc
                }
            }
            return {
                ...item
            }
        })
    },
    UpdateCampaignCountryItemRemove(state, index) {
        if (state.campaign.campaignCountries.length === 1) {
            state.campaign.campaignCountries = []
        } else {
            state.campaign.campaignCountries = state.campaign.campaignCountries.slice(0, index).
            concat(state.campaign.campaignCountries.slice(index + 1, state.campaign.campaignCountries.length))
        }
    }
}

const actions = {}

export default {
    namespaced: true,
    state: () => state,
    actions,
    mutations,
    getters,
};
