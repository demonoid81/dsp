import axios from "axios";

const state = {
    campaign: {},
    campaigns: []
}

const getters = {
    campaign: state => state.campaign,
    campaigns: state => state.campaigns,
    feedAudienceValue: state => {
        if (state.campaign.blacklist_feed) {
            return state.campaign.blacklist_feed
        }
        if (state.campaign.whitelist_feed) {
            return state.campaign.whitelist_feed
        }
        return []
    },
    sourceAudienceValue: state => {
        if (state.campaign.blacklist) {
            return state.campaign.blacklist
        }
        if (state.campaign.whitelist) {
            return state.campaign.whitelist
        }
        return []
    }
}

const mutations = {
    setCampaign(state, campaign) {
        state.campaign = campaign
    },
    addCampaignCountry(state) {
        state.campaign.countries.push({
            country: '',
            cpc: 0.0001
        })
    },
    updateCampaignCountryItemCountry(state, {country, index}) {
        state.campaign.countries = state.campaign.countries.map((item, itemIndex) => {
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
    updateCampaignCountryItemCPC(state, {cpc, index}) {
        state.campaign.countries = state.campaign.countries.map((item, itemIndex) => {
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
    updateCampaignCountryItemRemove(state, index) {
        if (state.campaign.countries.length === 1) {
            state.campaign.countries = []
        } else {
            state.campaign.countries = state.campaign.countries.slice(0, index).concat(state.campaign.countries.slice(index + 1, state.campaign.countries.length))
        }
    },
    setCampaignItemField(state, {value, name}) {
        state.campaign[name] = value
    },
    deleteCampaignItemField(state, name) {
        delete state.campaign[name]
    },
    addFeedAudience(state, {name, value}) {
        state.campaign[name].push(value)
    },
    setLimitsBudgetTotal(state, value) {
        let limit = {
            budget_total: value
        }
        if (state.campaign.limit) {
            limit = {
                ...state.campaign.limit,
                budget_total: value
            }
        }
        state.campaign = {
            ...state.campaign,
            limit
        }
    },
    setLimitsBudgetDaily(state, value) {
        let limit = {
            budget_daily: value
        }
        if (state.campaign.limit) {
            limit = {
                ...state.campaign.limit,
                budget_daily: value
            }
        }
        state.campaign = {
            ...state.campaign,
            limit
        }
    },
    setLimitsClickTotal(state, value) {
        let limit = {
            click_total: value
        }
        if (state.campaign.limit) {
            limit = {
                ...state.campaign.limit,
                click_total: value
            }
        }
        state.campaign = {
            ...state.campaign,
            limit
        }
    },
    setLimitsClickDaily(state, value) {
        let limit = {
            click_daily: value
        }
        if (state.campaign.limit) {
            limit = {
                ...state.campaign.limit,
                click_daily: value
            }
        }
        state.campaign = {
            ...state.campaign,
            limit
        }
    },
    setCampaigns (state, campaigns) {
        state.campaigns = campaigns
    },
    addCampaign (state) {
        let find = false
        state.dsp = state.campaigns.map(item => {
            if (item.id === state.campaign.id) {
                find = true
                return state.campaign
            }
            return item
        })
        if (!find) {
            state.campaigns.push(state.campaign)
        }
    },
}

const actions = {
    async setCampaign ({commit}) {
        const err = await axios.post('/api/campaigns', state.dspItem).then(
            () => {
                commit('addCampaign')
                return true
            }).catch( err => {
            console.log(err)
            return false;
        })
    },
    async getCampaigns ({commit}) {
        try {
            const { data } = await axios.get('/api/campaigns');
            commit('setCampaigns', data)
            return true
        } catch(err) {
            console.log(err)
            return false;
        }
    },
}

export default {
    namespaced: true,
    state: () => state,
    actions,
    mutations,
    getters,
};
