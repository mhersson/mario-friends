bundle: {
	apiVersion: "v1alpha1"
	name:       "mario-friends"
	instances: {
		troopa: {
			module: {
				url: "file://./app"
			}
			namespace: "mario-friends"
			values: {
				useIngress: true
				message:
					"""
						You've passed troopa! 🎉
						🎮 Let the games begin 🎮

						"""
				nextHopHostname: "bowser"
			}
		}
		bowser: {
			module: {
				url: "file://./app"
			}
			namespace: "mario-friends"
			values: {
				useIngress: true
				message:
					"""
						You've defeated bowser!
						You are doing great!😎

						"""
				nextHopHostname: "princess"
			}
		}
		princess: {
			module: {
				url: "file://./app"
			}
			namespace: "mario-friends"
			values: {
				useIngress: true
				message:
					"""
						You've saved the princess!
						🤩 🤩 🤩 🤩 🤩 🤩 🤩 🤩 🤩

						"""
			}
		}
	}
}
